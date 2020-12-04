#include <stdio.h>
#include <string.h>
#include <stdlib.h>

#define PLAYERS 426
#define MARBLES 7205800

struct {
    int map[MARBLES];
    int scores[PLAYERS];
    int size;
    int pos;
} game;

void print() {
    for(int i=0; i<game.size; i++) {
        if(i == game.pos) {
            printf("(%d) ", game.map[i]);
        } else {
            printf("%d ", game.map[i]);
        }
    }
    printf("\n");
}

void add(int value) {
    game.size++;

    for(int i=game.size-1; i>= game.pos + 1; i--) {
        game.map[i] = game.map[i-1];
    }

    game.map[game.pos] = value;
}

void rem() {
    for(int i=game.pos; i<game.size-1; i++) {
        game.map[i] = game.map[i+1];
    }

    game.size--;
}

void play(int player, int value) {
    if(value % 23 == 0) {
        // Score points!
        game.pos = (game.pos - 7) % game.size;
        int new_score = value + game.map[game.pos];
        game.scores[player] += new_score;

        // Remove the marble
        rem();
    } else {
        game.pos = ((game.pos + 1) % game.size) + 1;
        add(value);
    }
}

int main() {
    int player = 0;
    int winning_score = 0;

    game.map[0] = 0;
    game.size = 1;
    game.pos = 0;

    for(int i=1; i<=MARBLES; i++) {
        play(player, i);
        //print();

        player = (player + 1) % PLAYERS;
    }

    // Now find the best score
    for(int i=0; i<PLAYERS; i++) {
        if(game.scores[i] > winning_score) {
            winning_score = game.scores[i];
        }
    }
    printf("%d\n", winning_score);
}
