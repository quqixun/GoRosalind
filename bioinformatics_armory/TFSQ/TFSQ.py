from Bio import SeqIO
from io import StringIO

handle = StringIO('')
SeqIO.convert('rosalind_tfsq.txt', 'fastq', handle, 'fasta')

fasta = handle.getvalue()
with open('tfqs_output.txt', 'w') as f:
    f.write(fasta)
