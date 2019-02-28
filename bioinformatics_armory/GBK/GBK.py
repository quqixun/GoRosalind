from Bio import Entrez


genus, date_a, date_b = \
    map(lambda line: line.strip(),
        open('rosalind_gbk.txt').readlines())

term = '{}[Organism] AND ("{}"[PDAT] : "{}"[PDAT])'.format(
    genus, date_a, date_b)

Entrez.email = "quqixun@gmail.com"
handle = Entrez.esearch(db="nucleotide", term=term)
record = Entrez.read(handle)
handle.close()

print(int(record["Count"]))
