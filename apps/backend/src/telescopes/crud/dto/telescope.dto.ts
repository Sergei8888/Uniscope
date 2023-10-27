import {
    CameraDtoI,
    MountDtoI,
    MountType,
    OpticalSystemDtoI,
    OpticalSystemType,
    TelescopeDtoI,
    TelescopeStatus,
    TelescopeTarget,
} from '@uniscope/shared/nest';

import { UserDto } from '@/users/dto/user.dto';

export class TelescopeDto implements TelescopeDtoI {
    description?: string;
    id: number;
    imgUrl: string;
    latitude: number;
    longitude: number;
    name: string;
    status: TelescopeStatus;
    target: TelescopeTarget;
    zoom: number;
    camera: CameraDto;
    mount: MountDto;
    opticalSystem: OpticalSystemDto;
    owner: UserDto;
}

export class OpticalSystemDto implements OpticalSystemDtoI {
    aperture: number;
    focalLength: number;
    id: number;
    luminosity: number;
    manufacturer: string;
    model: string;
    type: OpticalSystemType;
}

export class MountDto implements MountDtoI {
    id: number;
    manufacturer: string;
    model: string;
    tier: number;
    type: MountType;
}

export class CameraDto implements CameraDtoI {
    height: number;
    id: number;
    manufacturer: string;
    model: string;
    pixelSize: number;
    quantumEfficiency: number;
    sensorSize: number;
    width: number;
}
