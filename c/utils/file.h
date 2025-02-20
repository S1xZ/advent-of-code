
#ifndef FILE_STRUCT_H
#define FILE_STRUCT_H

struct FileStruct {
    long* file_size;
    char* file_ptr;
};

#endif

#ifndef FILE_READER_H
#define FILE_READER_H

struct FileStruct readfile(char* filepath);

#endif
