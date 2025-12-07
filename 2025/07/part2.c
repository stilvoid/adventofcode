#include <stdio.h>
#include <string.h>

char map[256][256];
long long nums[256][256];

int rows = 0;
int cols = 0;

int main() {
    while(fgets(map[rows], 256, stdin)) {
        if(cols == 0) {
            cols = strlen(map[rows]);
        }
        rows++;
    };

    for(int row=1; row<rows; row++) {
        for(int col=0; col<cols; col++) {
            char cabove = map[row-1][col];
            long long nabove = nums[row-1][col];

            if(cabove == 'S') {
                nabove = 1;
            }

            if(nabove > 0) {
                char here = map[row][col];

                if(here == '^') {
                    if(col > 0) {
                        nums[row][col-1] += nabove;
                    }

                    if(col < cols-1) {
                        nums[row][col+1] += nabove;
                    }
                } else {
                    nums[row][col] += nabove;
                }
            }
        }
    }

    long long total = 0;
    for(int col=0; col<cols; col++) {
        total += nums[rows-1][col];
    }
    printf("Total: %lld\n", total);
}
