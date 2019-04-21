import subprocess
from Bio import SeqIO


def lalign36(a, b):
    command = ['lalign36', '-m', '10', a, b]
    output = open('align36_output.txt', 'w')
    subprocess.call(command, stdout=output)

    r = ''

    with open('align36_output.txt', 'r') as f:
        content = f.readlines()
    for i, line in enumerate(content):
        if 'lsw_sim' in line:
            sim = float(line.split()[-1])
            overlap = int(content[i + 1].split()[-1])
            if 32 <= overlap <= 40 and sim == 1.0:
                tr = ''
                for j in range(i + 9, len(content)):
                    if '>' in content[j]:
                        break
                    tr += content[j].strip()

                if r == '' or len(tr) < len(r):
                    r = tr

    return r


def hamming_diatance(seq1, seq2):
    return sum(c1 != c2 for c1, c2 in zip(seq1, seq2))


def count(seq, r):
    n = 0
    for i in range(len(seq) - len(r)):
        if hamming_diatance(seq[i:i + len(r)], r) < 4:
            n += 1
    return n


records = list(SeqIO.parse('rosalind_subo.txt', 'fasta'))
fastas = ['s.fasta', 't.fasta']
for i, record in enumerate(records):
    SeqIO.write(record, fastas[i], 'fasta')

r = lalign36('s.fasta', 't.fasta')
s = list(SeqIO.parse('s.fasta', 'fasta'))[0]
t = list(SeqIO.parse('t.fasta', 'fasta'))[0]

print(count(s.seq, r), count(t.seq, r))
