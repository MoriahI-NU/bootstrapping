#Code recreated from: 
library(mosaic)
library(jsonlite)

Mustangs <- read.csv("MustangPrice.csv")

#Sample mean
sample_mean <- mean( ~ price, data = Mustangs)

#Perform 2000 resampling trials
mustangs_price_boot <- do(2000) * mean( ~ price, data = resample(Mustangs))

#Calculate the 90% confidence interval using quantiles
quantile <- qdata( ~ mean, c(.05, .95), data = mustangs_price_boot)
confidence <- cdata(~ mean, 0.95, data = mustangs_price_boot)

#Calculate critical value
tstar <- qt(1-.95/2, df = 24)

#Margin of Error
t_margin <- tstar * sd( ~ mean, data = mustangs_price_boot)

result <- list(
  Sample_Mean = sample_mean,
  Confidence = confidence,
  tstar = tstar,
  t_marg = t_margin
)

cat(jsonlite::toJSON(result, pretty = TRUE), "\n")