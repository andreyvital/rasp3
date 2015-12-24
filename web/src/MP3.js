/* @flow */
export type MP3 = {
  file: string;
  size: number;
  id3?: {
    title?: string;
    artist?: string;
    album?: string;
    year?: number;
    track?: number;
  }
}
