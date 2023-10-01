export interface ScoreboardResponse {
    houses: HouseScore[];
    players: PlayerScore[];
}

export interface HouseScore {
    name: string;
    score: number;
}

export interface PlayerScore {
    name: string;
    house: string;
    score: number;
}