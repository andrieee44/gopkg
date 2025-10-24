#include <linux/ioctl.h>

unsigned int wrap_IOC(int dir, int type, int nr, int size) {
    return _IOC(dir, type, nr, size);
}
