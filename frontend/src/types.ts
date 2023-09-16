export interface Response {
    message: string
    services: Service[]
}

export interface Service {
    name: string
    ports: Port[]
    namespace: string
    type: string
    ip: string
    image: string | null
}

export interface Port {
    name: string
    port: number
}
