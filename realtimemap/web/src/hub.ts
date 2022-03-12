export interface PositionsDto {
  positions: PositionDto[];
}

export interface PositionDto {
  vehicle_id: string;
  longitude: number;
  latitude: number;
  heading: number;
  speed: number;
  doors_open: boolean;
}

export interface WebsocketProto {
  event_id: string;
  payload: string;
}

export interface GeoPoint {
  longitude: number;
  latitude: number;
}

export interface Viewport {
  southWest: GeoPoint;
  northEast: GeoPoint;
}

export interface UpdateViewport {
  viewport: Viewport;
}

export interface Notification {
  message: string;
}

export interface HubConnection {
  setViewport(swLng: number, swLat: number, neLng: number, neLat: number);

  onPositions(callback: (positions: PositionDto[]) => void);

  onNotification(callback: (notification: string) => void);

  disconnect(): Promise<void>;
}

function ByteBufferToObject(buff) {
  const enc = new TextDecoder('utf-8');
  const uint8Array = new Uint8Array(buff);
  const decodedString = enc.decode(uint8Array);
  // console.log(decodedString);
  return JSON.parse(decodedString);
}

function StringToArrayBuffer(str) {
  return new TextEncoder().encode(str);
}

class WebsocketConnect implements HubConnection {
  private connection: WebSocket;
  private onPositionsCallback?: (positions: PositionDto[]) => void;
  private onNotificationCallback?: (notification: string) => void;

  constructor() {
    const wsURL = `ws://localhost:7700/`;
    this.connection = new WebSocket(wsURL);
    this.connection.binaryType = 'arraybuffer';
    this.connection.onopen = this.onWebsocketOpen.bind(this);
    this.connection.onerror = this.onWebsocketError.bind(this);
    this.connection.onmessage = this.onWebsocketMessage.bind(this);
    this.connection.onclose = this.onWebsocketClose.bind(this);
  }

  onWebsocketOpen(event) {
    console.log('ws连接成功', event);
  }

  onWebsocketError(event) {
    console.error('ws错误', event);
  }

  onWebsocketMessage(event) {
    const proto = ByteBufferToObject(event.data);
    // console.log(proto);
    const data = JSON.parse(proto['payload']);
    // console.log(data);

    const eventId = proto['event_id'];
    if (eventId == 'positions') {
      if (this.onPositionsCallback != null) {
        this.onPositionsCallback(data);
      }
    } else if (eventId == 'notification') {
      if (this.onNotificationCallback != null) {
        this.onNotificationCallback(data);
      }
    }
  }

  onWebsocketClose(event) {
    console.log('ws连接关闭', event);
  }

  sendMessage(eventId, data) {
    const x: WebsocketProto = {
      event_id: eventId,
      payload: JSON.stringify(data),
    };
    const str = JSON.stringify(x);
    // console.log(str);
    this.connection.send(StringToArrayBuffer(str));
  }

  setViewport(swLng: number, swLat: number, neLng: number, neLat: number) {
    const x: Viewport = {
      southWest: {
        longitude: swLng,
        latitude: swLat,
      },
      northEast: {
        longitude: neLng,
        latitude: neLat,
      },
    };
    this.sendMessage('viewport', x);
  }

  onPositions(callback: (positions: PositionDto[]) => void) {
    this.onPositionsCallback = callback;
  }

  onNotification(callback: (notification: string) => void) {
    this.onNotificationCallback = callback;
  }

  async disconnect() {
    await this.connection.close(1000);
  }
}

export const connectToHub = new WebsocketConnect;
