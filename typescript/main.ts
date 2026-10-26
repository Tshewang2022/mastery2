import { mkdir, readFile } from 'node:fs/promises';
import { basename, dirname, join } from 'node:path';
import { fileURLToPath } from 'node:url';
import sharp from 'sharp';

export interface ImageInfo {
  format: string | undefined;
  width: number | undefined;
  height: number | undefined;
  path: string;
}

const imageDirectory = join(dirname(fileURLToPath(import.meta.url)), 'img');

export async function blobToImg(blob: Blob, fileName = 'image.png'): Promise<ImageInfo> {
  const outputPath = join(imageDirectory, basename(fileName));
  const imageData = Buffer.from(await blob.arrayBuffer());
  const image = await sharp(imageData).metadata();

  await mkdir(imageDirectory, { recursive: true });
  await sharp(imageData).toFile(outputPath);

  return {
    format: image.format,
    width: image.width,
    height: image.height,
    path: outputPath,
  };
}

async function main(): Promise<void> {
  const filePath = process.argv[2];
  if (!filePath) {
    throw new Error('Usage: node main.ts <image-file>');
  }

  const blob = new Blob([await readFile(filePath)]);
  console.log(blob, "where is the data");
  const image = await blobToImg(blob, filePath);
  console.log(JSON.stringify(image, null, 2));
}

main().catch((error: unknown) => {
  console.error(error instanceof Error ? error.message : error);
  process.exitCode = 1;
});
