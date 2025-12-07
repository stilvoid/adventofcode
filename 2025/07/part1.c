#include <stdio.h>
#include <string.h>

char map[256][256];

int main() {
    int rows = 0;
    int cols = 0;

    while(fgets(map[rows], 256, stdin)) {
        if(cols == 0) {
            cols = strlen(map[rows]);
        }
        rows++;
    };

    int splits = 0;
    
    for(int row=1; row<rows; row++) {
        for(int col=0; col< cols; col++) {
            char above = map[row-1][col];

            if(above == 'S' || above == '|') {
                char here = map[row][col];

                if(here == '^') {
                    if(col > 0) {
                        map[row][col-1] = '|';
                    }

                    if(col < cols-1) {
                        map[row][col+1] = '|';
                    }

                    splits++;
                } else {
                    map[row][col] = '|';
                }
            }
        }
    }

    printf("Splits: %d\n", splits);
}
