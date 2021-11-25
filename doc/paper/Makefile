PAPER = mhh22

FIGURES = 

BIB = deriv_parsing.bib

$(PAPER).pdf: $(PAPER).tex $(FIGURES)
	# pdflatex -interaction=nonstopmode $<
	# pdflatex -interaction=nonstopmode $<
	latex $<
	latex $<
	dvips $(PAPER).dvi
	ps2pdf $(PAPER).ps

rebuild-refs: $(PAPER).tex $(FIGURES) $(BIB)
	# pdflatex -interaction=nonstopmode $<
	# bibtex $(PAPER)
	# bibtex $(PAPER)
	# pdflatex -interaction=nonstopmode $<
	# pdflatex -interaction=nonstopmode $<
	latex $<
	latex $<
	bibtex $(PAPER)
	bibtex $(PAPER)
	latex $<
	latex $<
	dvips $(PAPER).dvi
	ps2pdf $(PAPER).ps

clean:
	-rm $(PAPER).pdf
	-rm $(PAPER).ps $(PAPER).dvi
	-rm $(PAPER).aux $(PAPER).bbl $(PAPER).blg $(PAPER).log $(PAPER).xcp