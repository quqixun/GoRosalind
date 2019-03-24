from Bio import SeqIO
from Bio.Alphabet import IUPAC


records = SeqIO.parse('rosalind_rvco.txt', 'fasta',
                      alphabet=IUPAC.unambiguous_dna)

n_match_rc = 0
for record in records:
    if record.seq == record.seq.reverse_complement():
        n_match_rc += 1
print(n_match_rc)
