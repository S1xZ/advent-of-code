#include "../../utils/file.h"
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

int main(int argc, char* argv[]) {
    int floor = 0;
    struct FileStruct dat = readfile("input.txt");
    for (long i = 0; i < *dat.file_size; i++) {
        if (dat.file_ptr[i] == '(') {
            floor += 1;
        }
        if (dat.file_ptr[i] == ')') {
            floor -= 1;
        }
    }
    printf("%d", floor);
    return EXIT_SUCCESS;
}
