
#include <ctype.h>
#include <stdbool.h>
#include <strings.h>
#include <stdio.h>

#define LENGTH 45


int main () {
        // TODO
    FILE *file = fopen("dictionaries/small", "r");
    bool result = true;
    char word[LENGTH + 2] = {'\0'};
    int index = 0, words = 0;

    if (file == NULL)
    {
        printf("Could not open %s.\n", "dictionaries/small");
        result = false;
        return result;
    }

    // read char by char until the end of the dictionary file
    char c;
    while (fgets(word, LENGTH + 2, file) != NULL)
    {
        words++;
        word[strlen(word) - 1] = '\0';
        for (int i = 0; i < LENGTH + 1; i++) {
            printf("%c",word[i]);
        }
        printf(" - strlen = %lu\n", strlen(word));
    }
    printf("words = %d\n", words);
}

