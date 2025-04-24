import torch
import numpy as np
from model.model import Generator  # stylegan2-pytorch
from torchvision.utils import save_image
from definitions import MODEL_PATH

device = torch.device('cuda')

with open(MODEL_PATH, 'rb') as f:
    g = Generator(1024, 512, 8).to(device)
    g.load_state_dict(torch.load(f)['g_ema'])
    g.eval()

# Load boundaries
boundaries = {
    'young': np.load('boundary/boundary_Young.npy'),
    'smiling': np.load('boundary/boundary_Smiling.npy'),
    'no_beard': np.load('boundary/boundary_No_Beard.npy'),
    # Добавь свои
}

# Convert boundaries by torch
for key in boundaries:
    boundaries[key] = torch.tensor(boundaries[key], dtype=torch.float32).to(device)

# Sample attributes
attributes = {
    'young': 1,
    'smiling': 0,
    'no_beard': 1,
}

# Create latent
z = torch.randn(1, 512).to(device)
w = g.style(z)

# Latent correction
for attr, flag in attributes.items():
    if flag == 1:
        w = w + boundaries[attr].unsqueeze(0) * 3.0  # 3.0 — сила влияния

# Image generation
with torch.no_grad():
    img, _ = g([w], input_is_latent=True)
    save_image((img.clamp(-1, 1) + 1) / 2, 'result.png')
