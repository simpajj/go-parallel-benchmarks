library(ggplot2)

go <- as.matrix(read.table("../seq-bubblesort.txt", col.names=c("type", "time")))
c <- as.matrix(read.table("../seq-bubblesort-c.txt", col.names=c("type", "time")))
data <- merge(go, c, all = TRUE)
png('seq-bubblesort.png')
barplot(go)
dev.off()