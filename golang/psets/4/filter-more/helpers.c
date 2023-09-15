#include "helpers.h"

void swap (RGBTRIPLE *a1, RGBTRIPLE *a2);

// Convert image to grayscale
void grayscale(int height, int width, RGBTRIPLE image[height][width])
{
    BYTE newVal = 0;
    for (int i = 0; i < height; i++)
    {
        for (int j = 0; j < width; j++)
        {
            newVal = (int)(image[i][j].rgbtBlue + image[i][j].rgbtGreen + image[i][j].rgbtRed) / 3;

            image[i][j].rgbtBlue = newVal;
            image[i][j].rgbtRed = newVal;
            image[i][j].rgbtGreen = newVal;
        }
    }
    return;
}

// Reflect image horizontally
void reflect(int height, int width, RGBTRIPLE image[height][width])
{
    for (int i = 0; i < height; i++)
    {
        for (int j = 0; j < width/2; j++)
        {
            swap(&image[i][j], &image[i][width - j]);
        }
    }
    return;
}

// Blur image
void blur(int height, int width, RGBTRIPLE image[height][width])
{
    int blue = 0, red = 0, green = 0;

    for (int i = 1; i < height - 1; i++)
    {
        for (int j = 1; j < width - 1; j++)
        {
            for (int k = -1; k < 2; k++){
                for (int f = -1; f < 2; f++) {
                    blue  += (int)image[i + k][j + f].rgbtBlue;
                    red   += (int)image[i + k][j + f].rgbtRed;
                    green += (int)image[i + k][j + f].rgbtGreen;
                }
            }
            image[i][j].rgbtBlue  = (BYTE)(blue/9);
            image[i][j].rgbtRed   = (BYTE)(red/9);
            image[i][j].rgbtGreen = (BYTE)(green/9);
            blue = 0; red = 0; green = 0;
        }
    }

    return;
}

// Detect edges
void edges(int height, int width, RGBTRIPLE image[height][width])
{
    
    return;
}

void swap (RGBTRIPLE *a1, RGBTRIPLE *a2){
    RGBTRIPLE temp = *a1;
    *a1 = *a2;
    *a2 = temp;
}