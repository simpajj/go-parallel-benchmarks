mydata <- as.matrix(read.table("../seq-bubblesort.txt"))
png('seq-bubblesort.png')
barplot(mydata, main="Bubblesort", ylab="Time", 
	names.arg=c("C", "Go"))
dev.off()