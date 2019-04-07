import subprocess
from Bio import SeqIO


with open('rosalind_filt.txt', 'r') as f:
    content = f.readlines()

q, p = content[0].strip().split(' ')
fastq = "".join(content[1:])
with open('filt_input.fastq', 'w') as f:
    f.write(fastq)

command = ['fastq_quality_filter', '-Q33', '-q', q, '-p', p,
           '-i', 'filt_input.fastq', '-o', 'filt_output.fastq']
subprocess.call(command)


records = SeqIO.parse('filt_output.fastq', 'fastq')
n = 0
for record in records:
    n += 1

print(n)
