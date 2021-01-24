/**
 * From: http://ccodearchive.net/info/array_size.html
*/

#include <stdio.h>

/**
 * ARRAY_SIZE - get the number of elements in a visible array
 * @arr: the array whose size you want.
 *
 * This does not work on pointers, or arrays declared as [], or
 * function parameters.  With correct compiler support, such usage
 * will cause a build error (see build_assert).
 */
#define ARRAY_SIZE(arr) (sizeof(arr) / sizeof((arr)[0]) + _array_size_chk(arr))

#if HAVE_BUILTIN_TYPES_COMPATIBLE_P && HAVE_TYPEOF
/* Two gcc extensions.
 * &a[0] degrades to a pointer: a different type from an array */
#define _array_size_chk(arr)						\
	BUILD_ASSERT_OR_ZERO(!__builtin_types_compatible_p(typeof(arr),	\
							typeof(&(arr)[0])))
#else
#define _array_size_chk(arr) 0
#endif


int main()
{
    char array[10] = {0};

    printf("array size is: %d\n", ARRAY_SIZE(array));

    return 0;
}
