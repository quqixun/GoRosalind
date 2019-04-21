import subprocess

in_file = 'rosalind_clus.txt'
command = ['clustalw', '-infile=' + in_file, '-align']
subprocess.call(command, stdout=subprocess.DEVNULL)

with open('rosalind_clus.txt', 'r') as f:
    content = f.readlines()
ids = [c.strip()[1:] for c in content if '>' in c]

with open('rosalind_clus.dnd', 'r') as f:
    diffs = f.readlines()

max_diff = 0
max_id = ''
for diff in diffs:
    for idd in ids:
        if idd in diff:
            diff = diff.strip()
            if diff[-1] == ',' or diff[-1] == ')':
                diff = diff[:-1]

            dnum = float(diff.split(':')[-1])
            if dnum > max_diff:
                max_id = idd
                max_diff = dnum

print(max_id)
