from Bio import SeqIO


with open('rosalind_bfil.txt', 'r') as f:
    content = f.readlines()

threshold = int(content[0].strip())
fastq = "".join(content[1:])
with open('bfil_input.fastq', 'w') as f:
    f.write(fastq)
records = list(SeqIO.parse('bfil_input.fastq', 'fastq'))

for record in records:

    quality = record.letter_annotations['phred_quality']
    qidx = [i for i, q in enumerate(quality) if q >= threshold]

    record.letter_annotations.clear()
    record.seq = record.seq[qidx[0]:qidx[-1] + 1]
    record.letter_annotations['phred_quality'] = quality[qidx[0]:qidx[-1] + 1]

SeqIO.write(records, 'bfil_output.fastq', 'fastq')
