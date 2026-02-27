#!/usr/bin/env bash
#SBATCH --job-name=network-test
#SBATCH --nodes=2
#SBATCH --partition=short
#SBATCH --time=00:30:00
#SBATCH --output=network-test-%j.out

NODES=$(scontrol show hostnames $SLURM_JOB_NODELIST)
NODE1=$(echo $NODES | awk '{print $1}')
NODE2=$(echo $NODES | awk '{print $2}')

echo "Testing: $NODE1 <-> $NODE2"

# Basic ping test
srun --nodes=1 --nodelist=$NODE1 ping -c 3 $NODE2
