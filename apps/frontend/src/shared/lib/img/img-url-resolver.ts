export class ImgUrlResolver {
    public static resolveBackendPath(url: string): string {
        if (import.meta.env.PROD) {
            return url;
        }

        return `http://localhost:9000${url}`;
    }
}
