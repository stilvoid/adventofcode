#include <stdio.h>
#include <stdlib.h>

typedef struct {
    long long x;
    long long y;
    long long z;
    int circuit;
} Box;

typedef struct {
    Box* a;
    Box* b;
    long long dist;
} Pair;

Box boxes[1000];
Pair pairs[1000 * 1000];
int circuits[1000];

int cmpPair(const void *a, const void *b) {
    const Pair *A = a;
    const Pair *B = b;

    if(A->dist < B->dist) {
        return -1;
    } else if(A->dist > B->dist) {
        return 1;
    }

    return 0;
}

int cmpInt(const void *a, const void *b) {
    const int *A = a;
    const int *B = b;

    return *B - *A;
}

int main() {
    int ncircuits = 0;
    int nboxes = 0;
    int npairs = 0;

    while(scanf("%lld,%lld,%lld", &boxes[nboxes].x, &boxes[nboxes].y, &boxes[nboxes].z) == 3) {
        boxes[nboxes].circuit = -1;
        nboxes++;
    }

    for(int i=0; i<nboxes-1; i++) {
        for(int j=i+1; j<nboxes; j++) {
            pairs[npairs].a = &boxes[i];
            pairs[npairs].b = &boxes[j];

            long long dx = boxes[i].x - boxes[j].x;
            long long dy = boxes[i].y - boxes[j].y;
            long long dz = boxes[i].z - boxes[j].z;

            pairs[npairs].dist = (dx*dx) + (dy*dy) + (dz*dz);

            npairs++;
        }
    }

    qsort(pairs, npairs, sizeof(pairs[0]), cmpPair);

    // Account for test mod
    int limit = 10;
    if(nboxes > 20) {
        limit = 1000;
    }

    // Connect pairs
    for(int i=0; i<limit; i++) {
        int aCirc = pairs[i].a->circuit;
        int bCirc = pairs[i].b->circuit;

        if(aCirc >= 0) {
            if(bCirc >= 0) {
                if(aCirc != bCirc) {
                    for(int j=0; j<nboxes; j++) {
                        if(boxes[j].circuit == bCirc) {
                            boxes[j].circuit = aCirc;
                        }
                    }
                }
            } else {
                pairs[i].b->circuit = aCirc;
            }
        } else if(pairs[i].b->circuit >= 0) {
            pairs[i].a->circuit = bCirc;
        } else {
            pairs[i].a->circuit = ncircuits;
            pairs[i].b->circuit = ncircuits;
            ncircuits++;
        }
    }

    // Count up the circuits
    for(int i=0; i<nboxes; i++) {
        if(boxes[i].circuit >= 0) {
            circuits[boxes[i].circuit]++;
        }
    }

    qsort(circuits, ncircuits, sizeof(int), cmpInt);

    int total = 0;
    for(int i=0; i<ncircuits; i++) {
        total += circuits[i];
    }

    printf("Answer %d\n", circuits[0] * circuits[1] * circuits[2]);

    return 0;
}
