# Bootstrapping and Simulation-Based Inference

### The Problem
The goal of this project was to select a task to perform in both R and Go and compare the two regarding their efficiency, memory requirements, and calculations.

I chose to perform a bootstrapping and simulation-based inference. For the Rscript in this repo, I recreated code from Daniel Kaplan, Nicholas J. Horton, and Randall Pruim (found here: https://cran.r-project.org/web/packages/mosaic/vignettes/Resampling.html). I used their proposed problem (referenced as a "Lock Problem" provided by Robin Lock et al. in 2011) as the backdrop for this experiment.

The Lock problem in question:  
"...construct a 90% confidence interval for the mean price of used Mustangs."

### The Data
MustangPrice.csv contains the input data for this experiment. It contains information relating to the age, price, and mileage of 25 different used Mustangs.  

The source for the Rscript did not provide a data file, so I collected information from www.autotrader.com and searched for used ford mustangs near me. The data in MustangPrice.csv was manually entered based on my findings.

### Setup and Packages
As mentioned previously, the Rscript in this repo (bootstrap.R) was based off of work done by Daniel Kaplan, Nicholas J. Horton, and Randall Pruim and can be found here: https://cran.r-project.org/web/packages/mosaic/vignettes/Resampling.html

The R package used for this experiment is the mosaic package: https://cran.r-project.org/web/packages/mosaic/readme/README.html  

For the Go code, I employed the "gonum.org/v1/gonum/stat/distuv" package for its distribution and mathematical functions. It was helpful in creating a function to find critical values. Other than this, I used very common packages like "math" or "math/rand" to create the rest of the helper functions (ie for bootstrapping, or finding confidence intervals) and ended up essentially creating my own Go package to carry out this project. You can find all these helper functions in the helper folder of this repo.

### Results
bootstrap.exe contains the final application for this experiment. Upon running, it will calculate the sample mean of MustangPrice.csv, the Confidence Intervals for the mean price of used Mustangs, the critical value (tStar), the margin of error (tMargin), and lastly the runtime. These values will be calculated and shown for both Rscript and Go.

Running the application on my system resulted in the following:  

---------R---------
Sample Mean: [31.964]
Confidence Intervals: [27.059300, 37.604500] 
tStar: [0.0634]
tMargin: [0.1767]
R elapsed time: 2.7199466s
---------Go---------
Sample Mean: 31.964000
Confidence Intervals (95%): [27.112000, 37.956000]
tStar: 0.063366
tMargin: 0.177238
Go elapsed time: 1.9064ms

As you can see, the calculations between these languages give very similar results. tStar and tMargin are both within rounding error of their counterparts. The confidence intervals vary the most - although not within rounding error their results have a differrence of about 0.35 or less which is noticeable but not overly concerning.

Obviously Go ran much faster, given that its runtime is recorded in milliseconds as opposed to seconds. If we look at the ratio R:Go (2719.9466ms:1.9064ms), Go runs at a rate of 1,427 times faster than R's runtime!

### Insight

Go can handle many projects with relative ease, and this one is no different. Despite the statistical requirements, it performed on par with R (which was built to handle complex stats) in terms of calculations, and way outperformed R in terms of runtime. The only real downside I found while carrying out this experiment was that Go did require more code. I would say that if concise or easy code is a priority then R might be the language to use because it has so many statistics-oriented packages available. However, if that is not your number one priority then I would say that (for the purposes of this experiment) Go outshines R in every other aspect.