all: beamer-tutorial17.pdf

%.pdf: %.tex
	pdflatex -shell-escape $<

.PHONY: clean
clean:
	rm *.log *.aux *.nav *.out *.snm *.toc *.vrb

