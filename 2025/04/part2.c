#include <stdio.h>
#include <stdlib.h>
#include <strings.h>

char get(char* data, int x, int y, int w, int h) {
    if(x<0 || y<0 || x>=w || y>=h) {
        return '.';
    }

    return data[y*w+x];
}

int set(char* data, int x, int y, int w, int h, char c) {
    if(x<0 || y<0 || x>=w || y>=h) {
        return 0;
    }

    data[y*w+x] = c;

    return 1;
}

int count(char* data, int x, int y, int w, int h) {
    if(get(data, x, y, w, h) == '.') {
        return -1;
    }

    int count = 0;

    for(int i=-1; i<=1; i++) {
        for(int j=-1; j<=1; j++) {
            if(i != 0 || j != 0) {
                if(get(data, x+i, y+j, w, h) == '@') {
                    count++;
                }
            }
        }
    }

    return count;
}

int step(char* data, int w, int h) {
    int total = 0;
    int c = 0;

    char* newdata = calloc(w * h, sizeof(char));
    memcpy(newdata, data, w * h);

    for(int y=0; y<h; y++) {
        for(int x=0; x<w; x++) {
            c = count(data, x, y, w, h);
            if(c >= 0 && c<4) {
                total++;
                set(newdata, x, y, w, h, '.');
            }
        }
    }

    memcpy(data, newdata, w*h);

    return total;
}

int main(int argc, char* argv[]) {
    int w = 0, h = 0;

    char line[256];

    char* data = calloc(20 * 1024, sizeof(char));

    while(scanf("%s", line) == 1) {
        if(w==0) {
            w=strlen(line);
        }

        strcat(data, line);
        h++;
    }

    int count = -1;
    int total = 0;
    while(count != 0) {
        count = step(data, w, h);
        total += count;
    }
    printf("Total: %d\n", total);

    return 0;
}
