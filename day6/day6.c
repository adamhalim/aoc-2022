#include <stdio.h>
#include <string.h>

int lookForMarker(FILE *fp, char* line, int markerSize);
int checkSequence(char c, char *sequence, int *sequenceIndex, int markerSize);

char firstSequence[4];
char secondSequence[14];
int firstSeqIdx;
int secondSeqIdx;

#define STR_LEN 4096

int main() {
    FILE *fp = fopen("input.txt", "r");
    char line[STR_LEN];

    int firstMarker = lookForMarker(fp, line, 4);
    fseek(fp, 0, SEEK_SET);
    int secondMarker = lookForMarker(fp, line, 14);

    printf("First marker after character %d.\n", firstMarker);
    printf("Second marker after character %d.\n", secondMarker);
}

int lookForMarker(FILE *fp, char* line, int markerSize) {
    fgets(line, STR_LEN, fp);
    int firstMarker = 0;
    for(int i = 0; i < STR_LEN; i++) {
        if (!firstMarker) {
            int markerFound = checkSequence(line[i], firstSequence, &firstSeqIdx, markerSize);
            if (markerFound == markerSize) 
                firstMarker = i;
            else
                i -= markerFound;
        }
    }
    return firstMarker;
}

int checkSequence(char c, char* sequence, int *sequenceIndex, int markerSize) {
    char *charInSequence = strchr(sequence, c);
    if (!charInSequence) {
        if (*sequenceIndex == markerSize) return markerSize;
        sequence[(*sequenceIndex)++] = c;
        return 0;
    }

    // Reset sequence and try again
    memset(sequence, 0, markerSize);
    int idx = *sequenceIndex;
    *sequenceIndex = 0;
    return idx;
}