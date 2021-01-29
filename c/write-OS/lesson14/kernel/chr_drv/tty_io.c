#include <linux/kernel.h>
#include <linux/head.h>
#include <linux/sched.h>
#include <asm/system.h>
#include <asm/segment.h>
#include <asm/io.h>
#include <serial_debug.h>
#include <linux/tty.h>

#define DEBUG

void con_init();
extern void con_write(struct tty_struct *tty);

// 这里我们用 C99 的 dot initializer 初始化
struct tty_struct tty_table[] = {
    {
        .pgrp = 0,
        .write = NULL,
        .flags = 0 | TTY_ECHO,
        .read_q = {
            .head = 0,
            .tail = 0,
        },
        .write_q = {
            .head = 0,
            .tail = 0,
        },
        .buffer = {
            .head = 0,
            .tail = 0
        }
    },
};

// struct tty_queue read_q;
void tty_init(void) {
    tty_table[0].write = con_write;
    con_init();
}

// copy_to_buffer handle the keyboard input
// to tty echo char
void copy_to_buffer(struct tty_struct *tty) {
    char ch;
    struct tty_queue *read_q= &tty->read_q;
    struct tty_queue *buffer= &tty->buffer;
    while(!tty_isempty_q(&tty->read_q)) {
        ch = tty_pop_q(read_q);
        // Judge if it's a normal character
        // if (ch < 32) {
        //     // This is control character
        //     continue;
        // }
        switch(ch) {
            case '\b':
                // This is backspace char
                if (!tty_isempty_q(buffer)) {
                    if(tty_queue_tail(buffer) == '\n')  // \n 不能被清除掉
                       continue ;
                    buffer->tail = (buffer->tail - 1) % TTY_BUF_SIZE;
                } else {
                    continue;
                }
                break;
            case -1:
            case '\n':
                s_printk("Wakeup!\n");
                wake_up(&tty_table[0].buffer.wait_proc);
                break;
                // EOF
            default:
                if (!tty_isfull_q(buffer)) {
                    tty_push_q(buffer, ch);
                } else {
                    // here we need to sleep until the queue
                    // is not full
                }
                break;
        }
        if (tty->flags | TTY_ECHO) {
            tty_push_q(&tty->write_q, ch);
            tty->write(tty);
        }
    }
    return ;
}

void tty_write(struct tty_struct* tty) {
    tty->write(tty);
    return ;
}

// 队列为空时，进程进入睡眠状态, 可被中断唤醒
// TODO: write sleep_if_empty
// static void sleep_if_empty(struct tty_queue *q) {
//     cli();
//     while(!current->signal && tty_isempty_q(q))
//         interruptible_sleep_on(&q->wait_proc);
//     sti();
// }

// static void sleep_if_full(struct tty_queue *q) {
//     cli();
//     while(!current->signal && tty_isfull_q(q))
//         interruptible_sleep_on(&q->wait_proc);
//     sti();
// }

// tty_read 函数从 tty->buffer 中读取内容
// 当没有读到足够的字符且 buffer 为空的时候
// sleep
// nr = -1 means only stop at \n
int tty_read(int channel, char *buf, int nr) {
    int len = 0;
    char ch;
    char *p = buf;
    // unsigned int i;
    // return from the tty_read when recv a \n or EOF
    // char done = 0;
    //unsigned int headptr = 0;
    // Sanity check
#ifdef DEBUG
    s_printk("buf addr = 0x%x\n", buf);
#endif
    if (channel > 2 || channel < 0 || nr < 0)
        return -1;
    struct tty_struct *tty = tty_table + channel;
    // Sleep until the queue is not empty
    // sleep_if_empty(&tty->buffer);
    // Only we get an \n, we start to process
    interruptible_sleep_on(&tty->buffer.wait_proc);

    while (nr > 0) {
        ch = tty_pop_q(&tty->buffer);
        // TODO: Change -1 to EOF
        if (ch == '\n' || ch == -1) {
            // \n Or EOF, we are done
            // done = 1;
            // Now we can just stop and pop all to buffer
            break;
        }
        /* if (ch == CTRL_INT) return -1; */
        put_fs_byte(ch, p);
        *p++ = ch;
#ifdef DEBUG
        s_printk("Buf = %s\n", buf);
#endif
        len++;
        nr--;
    }
    // while (!done) {
    //     sleep_if_empty(&tty->buffer);
    //     for (i = tty->buffer.head; i != tty->buffer.tail; i = (i + 1) % TTY_BUF_SIZE) {
    //         if (tty->buffer.buf[i] == '\n' || tty->buffer.buf[i] == -1) {
    //             done = 1;
    //             break;
    //         }
    //     }
    // }
    return len;
}
