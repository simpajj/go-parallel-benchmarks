/* http://stackoverflow.com/questions/13719965/mandelbrot-set-in-openmp */

#include <stdio.h>
#include <stdlib.h>
#include <math.h>
#include <string.h>
#include <unistd.h>
#include <stdint.h>
#include <omp.h>

static const int ImageWidth  = 1000;
static const int ImageHeight = 1000;
static const double CxMin    = -2.5;
static const double CxMax    = 1.5;
static const double CyMin    = -2.0;
static const double CyMax    = 2.0;

int main(int argc, char *argv[])
{
	omp_set_dynamic(0);
	omp_set_num_threads(atoi(argv[1])); // Number of threads
	const int N = atoi(argv[2]); // Number of iterations
	int i;

  double PixelWidth=(CxMax-CxMin)/ImageWidth;
  double PixelHeight=(CyMax-CyMin)/ImageHeight;

  const int MaxColorComponentValue=1<<8;
  typedef unsigned char pixel_t[3];
  pixel_t *pixels = malloc(sizeof(pixel_t)*ImageHeight*ImageWidth);

  for (i = 0; i <= N; i++) {
	  #pragma omp parallel shared(pixels)
	  {
	    int iY;
	    #pragma omp for
	    for(iY=0; iY<ImageHeight; iY++) {
	      double Cy=CyMin + iY*PixelHeight;
	      if (fabs(Cy)< PixelHeight/2) {
	        Cy=0.0;
	      }
	      int iX;
	      for(iX=0; iX<ImageWidth; iX++) {
	        double Cx=CxMin + iX*PixelWidth;
	        double Zx=0.0;
	        double Zy=0.0;
	        double Zx2=Zx*Zx;
	        double Zy2=Zy*Zy;
	        int Iteration;
	        const int IterationMax=150;
	        const double Bailout=2;
	        const double Circle_Radius=Bailout*Bailout;

	        for (Iteration=0; Iteration<IterationMax && ((Zx2+Zy2)<Circle_Radius); Iteration++) { // 
	            Zy=2*Zx*Zy + Cy;
	            Zx=Zx2-Zy2 + Cx;
	            Zx2=Zx*Zx;
	            Zy2=Zy*Zy;
	        };

	        if (Iteration==IterationMax) {
	            pixels[iY*ImageWidth + iX][0] = 0;
	            pixels[iY*ImageWidth + iX][1] = 0;
	            pixels[iY*ImageWidth + iX][2] = 0;
	        }
	        else {
	          pixels[iY*ImageWidth + iX][0] = ((double)(Iteration-log2(log2(sqrt(Zx2+Zy2))))/IterationMax) * MaxColorComponentValue;
	          pixels[iY*ImageWidth + iX][1] = 0;
	          pixels[iY*ImageWidth + iX][2] = 0;
	        }
	      }
	    }
	  }
	}
  return 0;
}