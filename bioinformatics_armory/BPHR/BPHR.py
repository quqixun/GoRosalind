import numpy as np

from Bio import SeqIO


with open('rosalind_bphr.txt', 'r') as f:
    content = f.readlines()

threshold = int(content[0].strip())
fastq = "".join(content[1:])
with open('bphr_input.fastq', 'w') as f:
    f.write(fastq)
records = list(SeqIO.parse('bphr_input.fastq', 'fastq'))

qualities = []
for record in records:
    qualities.append(record.letter_annotations['phred_quality'])

mean_quality = np.mean(qualities, axis=0)
n_below_threshold = len(np.where(mean_quality < threshold)[0])
print(n_below_threshold)
