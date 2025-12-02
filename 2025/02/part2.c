#include <stdio.h>
#include <string.h>

int check(char* subject) {
    for(int len=1; len<strlen(subject); len++) {
        if(strlen(subject) % len == 0) {
            int size = strlen(subject) / len;
            int match = -1;

            for(int start=len; start<strlen(subject); start+=len) {
                if(strncmp(subject, subject+start, len) != 0) {
                    match = 0;
                    break;
                }
            }

            if(match) {
                return match;
            }
        }
    }

    return 0;
}

int main(int argc, char* argv[]) {
    if(argc!=2) {
        fprintf(stderr, "Need INPUT\n");
        return 1;
    }

    FILE* f = fopen(argv[1], "r");

    int max = 256;
    int n;
    char buf[max];

    unsigned long result = 0;

    while(!feof(f)) {
        unsigned long start, end;

        fscanf(f, "%lu-%lu", &start, &end);

        for(unsigned long i=start; i<=end; i++) {
            n = snprintf(buf, max, "%lu", i);
            if(n>=max) {
                fprintf(stderr, "number too long!\n");
                return 1;
            }

            if(check(buf)) {
                result += i;
            }
        }

        if(!feof(f)) {
            char c = fgetc(f);
            if(c != ',') {
                break;
            }
        }
    }

    printf("RESULT: %lu\n", result);

    fclose(f);

    return 0;
}
