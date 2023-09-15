// Implements a dictionary's functionality

#include <ctype.h>
#include <stdbool.h>
#include <strings.h>
#include <stdio.h>
#include <stdlib.h>

#include "dictionary.h"

void chainNodeToHashTable(const char *item);

// Represents a node in a hash table
typedef struct node
{
    char word[LENGTH + 1];
    struct node *next;
}
node;

// TODO: Choose number of buckets in hash table
// this value will be changed along with hash function
const unsigned int N = 187751;

// Hash table
node *table[N] = {NULL};
// size of dictionary
static unsigned int sizeOfDictionary = 0;

// Hashes word to a number
unsigned int hash(const char *word)
{
    // TODO: Improve this hash function
    // return (toupper(word[0]) - 'A');
    int hash = 0;
    for (int i = 0; i <= strlen(word); i++)
    {
        // Convert the character to lowercase using the tolower() function.
        char c = tolower(word[i]);
        
        // Use the ASCII value of the character as part of the hash calculation.
        // Multiply the current hash value by 31 and add the ASCII value of the character.
        hash = (hash * 31 + c) % N; // N is the size of your hash table
    }
    return hash;
}

// Returns true if word is in dictionary, else false
// TODO: Your spell checker must not leak any memory. Be sure to check for leaks with valgrind.
bool check(const char *word)
{
    // strcasecmp in strings.h to compare two strings case-insensitively,
    // TODO
    node *hold = table[hash(word)];

    while (hold != NULL)
    {
        if ((strcasecmp(hold->word, word) == 0))
        {
            return true;
        }
        hold = hold->next;
    }

    return false;
}

// chainNodeToHashTable
void chainNodeToHashTable(const char *item) {

    node *newNode = (node*)malloc(sizeof(node));
    unsigned int hashIndex = hash(item);

    strcpy(newNode->word, item);
    // prepend newnode to the list
    newNode->next = table[hashIndex];
    table[hashIndex] = newNode;
}


// Loads dictionary into memory, returning true if successful, else false
bool load(const char *dictionary)
{
    // TODO
    FILE *file = fopen(dictionary, "r");
    bool result = true;
    char word[LENGTH + 2]; // include the line feed of each line in dictionary
    int index = 0, words = 0;

    if (file == NULL)
    {
        printf("Could not open %s.\n", dictionary);
        result = false;
        return result;
    }

    // read char by char until the end of the dictionary file
    char c;
    while (fgets(word, LENGTH + 2, file) != NULL)
    {
        words++;
        //fgets will get the '\n' along
        word[strlen(word) - 1] = '\0';
        chainNodeToHashTable(word);
    }

    // Check whether there was an error
    if (ferror(file))
    {
        fclose(file);
        printf("Error reading %s.\n", dictionary);
        unload();
        result = false;
    }
    else
    {
        sizeOfDictionary = words;
    }

    return result;
}

// Returns number of words in dictionary if loaded, else 0 if not yet loaded
unsigned int size(void)
{
    // TODO
    return sizeOfDictionary;
}

// Unloads dictionary from memory, returning true if successful, else false
// Free the memory
bool unload(void)
{
    // TODO
    bool result = true;
    node *nodePtr = NULL;
    node *list = NULL;

    for (int i = 0; i < N; i++)
    {
        list = table[i];
        while (list != NULL)
        {
            nodePtr = table[i]->next;
            free(list);
            list = nodePtr;
        }
    }

    return result;
}
