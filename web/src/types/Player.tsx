import { atom } from 'recoil';

export interface Player {
  id: number;
  name: string;
  wandID: number;
  house: string;
  progress: PlayerProgress;
}

export interface PlayerProgress {
  sortingHat: boolean;
  pensieve: boolean;
}

export const PlayerState = atom<Player | undefined>({
    key: 'PlayerState',
    default: undefined
  })

export default PlayerState