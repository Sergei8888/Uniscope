import { UserDtoI } from './user';

export interface TelescopeDtoI {
    id: number;
    name: string;
    description?: string;
    imgUrl: string;
    latitude: number;
    longitude: number;
    zoom: number;
    target: TelescopeTarget;
    status: TelescopeStatus;
    opticalSystem: OpticalSystemDtoI;
    camera: CameraDtoI;
    mount: MountDtoI;
    owner: UserDtoI;
}

export enum TelescopeTarget {
    Sun = 'sun',
    NightSky = 'nightsky',
}

export enum TelescopeStatus {
    Online = 'online',
    Offline = 'offline',
    Booked = 'booked',
    Busy = 'busy',
}

export interface OpticalSystemDtoI {
    id: number;
    aperture: number;
    manufacturer: string;
    model: string;
    type: OpticalSystemType;
    luminosity: number;
    focalLength: number;
}

export enum OpticalSystemType {
    Refractor = 'refractor',
    Reflector = 'reflector',
    ShmidtCassegrain = 'shmidtCassegrain',
}

export interface CameraDtoI {
    id: number;
    manufacturer: string;
    model: string;
    sensorSize: number;
    pixelSize: number;
    quantumEfficiency: number;
    width: number;
    height: number;
}

export interface MountDtoI {
    id: number;
    manufacturer: string;
    model: string;
    tier: number;
    type: MountType;
}

export enum MountType {
    Azimutal = 'azimutal',
    Equatorial = 'equatorial',
}
