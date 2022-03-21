import { Reader, Writer } from "protobufjs/minimal";
export declare const protobufPackage = "christine713.getweather.getweather";
export interface MsgQueryWether {
    creator: string;
    title: string;
    body: string;
}
export interface MsgQueryWetherResponse {
}
export declare const MsgQueryWether: {
    encode(message: MsgQueryWether, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgQueryWether;
    fromJSON(object: any): MsgQueryWether;
    toJSON(message: MsgQueryWether): unknown;
    fromPartial(object: DeepPartial<MsgQueryWether>): MsgQueryWether;
};
export declare const MsgQueryWetherResponse: {
    encode(_: MsgQueryWetherResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgQueryWetherResponse;
    fromJSON(_: any): MsgQueryWetherResponse;
    toJSON(_: MsgQueryWetherResponse): unknown;
    fromPartial(_: DeepPartial<MsgQueryWetherResponse>): MsgQueryWetherResponse;
};
/** Msg defines the Msg service. */
export interface Msg {
    /** this line is used by starport scaffolding # proto/tx/rpc */
    QueryWether(request: MsgQueryWether): Promise<MsgQueryWetherResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    QueryWether(request: MsgQueryWether): Promise<MsgQueryWetherResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
