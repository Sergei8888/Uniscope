import { IsLatitude, IsLongitude } from 'class-validator';

export class LocationDto {
    /**
     * @example 111.12
     */
    @IsLatitude()
    public lat: number;

    /**
     * @example 111.12
     */
    @IsLongitude()
    public lng: number;
}
