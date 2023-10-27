import { OpticalSystemType, TelescopeTarget } from '@uniscope/shared/nest';

export class TelescopeLeadDto {
    public telescopeName: string;

    public telescopeDescription?: string;

    public locationLat: number;

    public locationLng: number;

    public telescopeTarget: TelescopeTarget;

    public cameraManufacturer: string;

    public cameraModel: string;

    public cameraResolutionWidth: number;

    public cameraResolutionHeight: number;

    public cameraSensorSize?: number;

    public cameraPixelSize?: number;

    public cameraQuantumEfficiency?: number;

    public opticalSystemManufacturer: string;

    public opticalSystemModel: string;

    public opticalSystemAperture?: number;

    public opticalSystemType?: OpticalSystemType;

    public opticalSystemFocalLength?: number;

    public opticalSystemLuminosity?: number;

    public telescopeZoom?: number;
}
