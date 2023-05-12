#include <stdio.h>

struct rect {
    float x;
    float y;
};

struct rect *new_rect(float x, float y) {
    struct rect r = {.x = x, .y = y};
    return &r;
}


int main(int argc, char* argv[]) {
    struct rect *rp = new_rect(1.2, 3.4);
    printf("address = %p\n", rp);
    printf(" (x, y) = (%.2f, %.2f)\n", rp->x, rp->y);
    return 0;
}
