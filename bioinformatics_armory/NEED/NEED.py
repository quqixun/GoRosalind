# Firstly, install Jemboss using command:
# sudo apt-get install jemboss

from Bio import Entrez, SeqIO
from Bio.Emboss.Applications import NeedleCommandline


Entrez.email = 'quqixun@gmail.com'

with open('rosalind_need.txt', 'r') as f:
    ids = f.read().split()

handle = Entrez.efetch(db='nucleotide', id=ids, rettype='fasta')
records = list(SeqIO.parse(handle, 'fasta'))

for i, record in enumerate(records):
    with open(ids[i], 'w') as f:
        SeqIO.write(record, f, 'fasta')

# Following step can be done in command line
# using 'needle'
needle_cl = NeedleCommandline()
needle_cl.asequence = ids[0]
needle_cl.bsequence = ids[1]
needle_cl.outfile = 'need_output.txt'

needle_cl.gapopen = 10
needle_cl.gapextend = 1
needle_cl.endopen = 10
needle_cl.endextend = 1
needle_cl.endweight = True

needle_cl()

# Score
with open('need_output.txt', 'r') as f:
    content = f.readlines()

for line in content:
    if 'Score' in line:
        score = float(line.split()[-1])
        print(score)
        break
