import warnings

from Bio.Alphabet import IUPAC
from Bio.Seq import Seq, translate


warnings.filterwarnings('ignore')


with open('rosalind_orfr.txt', 'r') as f:
    seq_text = f.read().strip()
seq = Seq(seq_text, IUPAC.unambiguous_dna)

DNAs = [seq, seq.reverse_complement()]
longest_protein = ''

for i in range(3):
    for DNA in DNAs:
        protein = translate(DNA[i:], to_stop=True)
        if len(longest_protein) < len(protein):
            longest_protein = protein

print(longest_protein)
