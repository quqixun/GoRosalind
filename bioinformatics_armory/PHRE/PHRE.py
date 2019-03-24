from Bio import SeqIO


with open('rosalind_phre.txt', 'r') as f:
    content = f.readlines()

threshold = int(content[0].strip())
fastq = "".join(content[1:])
with open('phre_fastq.txt', 'w') as f:
    f.write(fastq)
records = SeqIO.parse('phre_fastq.txt', 'fastq')

n_below_threshold = 0
for record in records:
    quality = record.letter_annotations['phred_quality']
    if float(sum(quality)) / len(quality) < threshold:
        n_below_threshold += 1

print(n_below_threshold)
