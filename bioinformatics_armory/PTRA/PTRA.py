from Bio.Seq import translate


text = []
with open('rosalind_ptra.txt', 'r') as f:
    for t in f:
        text.append(t.strip())
DNA, protein = text

igc = 0
NCBI_identifiers = [1, 2, 3, 4, 5, 6, 9,
                    10, 11, 12, 13, 14, 15]
for ni in NCBI_identifiers:
    trans_protein = translate(DNA, table=ni, to_stop=True)
    if trans_protein == protein:
        igc = ni
        break

print(igc)
