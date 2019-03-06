from Bio import SeqIO
from Bio import Entrez

with open('rosalind_frmt.txt', 'r') as f:
    ids = f.read().split()

Entrez.email = "quqixun@gmail.com"
handle = Entrez.efetch(db="nucleotide", id=ids, rettype="fasta")
records = list(SeqIO.parse(handle, "fasta"))
handle.close()

rlens = [len(r.seq) for r in records]
min_len_record = records[rlens.index(min(rlens))]
fasta_output = min_len_record.format('fasta')

with open('rosalind_frmt_res.txt', 'w') as f:
    f.write(fasta_output)
