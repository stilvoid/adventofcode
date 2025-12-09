#include <stdio.h>
#include <stdlib.h>

typedef struct {
    long long x;
    long long y;
} Point;

typedef struct {
    Point* A;
    Point* B;
    long long area;
} Rect;

#define MAXPOINTS 500

Point points[MAXPOINTS];
Rect rects[MAXPOINTS * MAXPOINTS];

int npoints = 0;
int nrects = 0;
int size=0;

char *grid;

long long xs[MAXPOINTS];
long long ys[MAXPOINTS];

int cmpRects(const void *A, const void *B) {
    const Rect *a = A;
    const Rect *b = B;

    if(a->area < b->area) {
        return 1;
    } else if(a->area > b->area) {
        return -1;
    }
    return 0;
}

int cmpint(const void *A, const void *B) {
    const long long *a = A;
    const long long *b = B;

    if(*a < *b) {
        return -1;
    } else if (*a > *b) {
        return 1;
    }
    return 0;
}

int find_x(long long x) {
    for(int i=0; i<size; i++) {
        if(xs[i] == x) {
            return i + 1;
        }
    }
    fprintf(stderr, "Error finding x: %lld\n", x);
    exit(1);
    return -1;
}

int find_y(long long y) {
    for(int i=0; i<size; i++) {
        if(ys[i] == y) {
            return i + 1;
        }
    }
    fprintf(stderr, "Error finding y: %lld\n", y);
    exit(1);
    return -1;
}

long long area(Point *A, Point *B) {
    long long dx = A->x - B->x;
    long long dy = A->y - B->y;

    if(dx<0) {
        dx=-dx;
    }

    if(dy<0) {
        dy=-dy;
    }

    return (dx+1)*(dy+1);
}

void draw_line(Point* a, Point* b) {
    int x1 = find_x(a->x);
    int y1 = find_y(a->y);
    int x2 = find_x(b->x);
    int y2 = find_y(b->y);

    if(x1>x2) {
        int bx = x1;
        x1 = x2;
        x2 = bx;
    }

    if(y1>y2) {
        int by=y1;
        y1 = y2;
        y2 = by;
    }

    for(int y=y1; y<=y2; y++) {
        for(int x=x1; x<=x2; x++) {
            grid[y * size + x] = '#';
        }
    }
}

void fill(int x, int y) {
    char c = grid[y*size+x];

    if(c == ' ') {
        grid[y*size+x] = '.';
        if(x<size-1) {
            fill(x+1, y);
        }
        if(y<size-1) {
            fill(x, y+1);
        }
        if(x>0) {
            fill(x-1, y);
        }
        if(y>0) {
            fill(x, y-1);
        }
    }
}

int check_rect(Rect* r) {
    int x1 = find_x(r->A->x);
    int y1 = find_y(r->A->y);
    int x2 = find_x(r->B->x);
    int y2 = find_y(r->B->y);

    if(x1>x2) {
        int bx = x1;
        x1 = x2;
        x2 = bx;
    }

    if(y1>y2) {
        int by = y1;
        y1 = y2;
        y2 = by;
    }

    for(int x=x1;x<=x2;x++) {
        for(int y=y1;y<=y2;y++) {
            if(grid[y*size+x] == '.') {
                return 0;
            }
        }
    }
    
    return 1;
}

int main() {
    while(scanf("%lld,%lld", &points[npoints].x, &points[npoints].y) == 2) {
        npoints++;
    }

    for(int i=0; i<npoints; i++) {
        xs[i] = points[i].x;
        ys[i] = points[i].y;
    }

    qsort(xs, npoints, sizeof(xs[0]), cmpint);
    qsort(ys, npoints, sizeof(ys[0]), cmpint);

    size = npoints + 2;

    grid = malloc(size * size);
    for(int i=0; i<size * size; i++) {
        grid[i] = ' ';
    }

    for(int i=0; i<npoints; i++) {
        int j = (i + 1) % npoints;
        draw_line(&points[i], &points[j]);
    }

    fill(0, 0);

    // Find all rects
    for(int i=0; i<npoints-1; i++) {
        for(int j=i+1; j<npoints; j++) {
            Point *A = &points[i];
            Point *B = &points[j];

            rects[nrects].A = A;
            rects[nrects].B = B;

            if(check_rect(&rects[nrects])) {
                rects[nrects].area = area(A, B);
            }

            nrects++;
        }
    }

    qsort(rects, nrects, sizeof(rects[0]), cmpRects);

    printf("Biggest: %lld\n", rects[0].area);
}
