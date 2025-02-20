#include "file.h"
#include <stdio.h>
#include <stdlib.h>

struct FileStruct readfile(char* filepath) {
    char* buffer = 0;
    long length;
    FILE* f = fopen(filepath, "rb");

    if (f) {
        fseek(f, 0, SEEK_END);
        length = ftell(f);
        fseek(f, 0, SEEK_SET);
        buffer = malloc(length);
        if (buffer) {
            fread(buffer, 1, length, f);
        }
        fclose(f);
    }
    struct FileStruct file = {&length, buffer};
    if (buffer) {
        return file;
    }
    exit(0);
}
