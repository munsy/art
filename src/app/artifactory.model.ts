// this file was automatically generated, DO NOT EDIT

// helpers
const maxUnixTSInSeconds = 9999999999;

function ParseDate(d: Date | number | string): Date {
        if (d instanceof Date) return d;
        if (typeof d === 'number') {
                if (d > maxUnixTSInSeconds) return new Date(d);
                return new Date(d * 1000); // go ts
        }
        return new Date(d);
}

function ParseNumber(v: number | string, isInt = false): number {
        if (!v) return 0;
        if (typeof v === 'number') return v;
        return (isInt ? parseInt(v) : parseFloat(v)) || 0;
}

function FromArray<T>(Ctor: { new(v: any): T }, data?: any[] | any, def = null): T[] | null {
        if (!data || !Object.keys(data).length) return def;
        const d = Array.isArray(data) ? data : [data];
        return d.map((v: any) => new Ctor(v));
}

function ToObject(o: any, typeOrCfg: any = {}, child = false): any {
        if (!o) return null;
        if (typeof o.toObject === 'function' && child) return o.toObject();

        switch (typeof o) {
                case 'string':
                        return typeOrCfg === 'number' ? ParseNumber(o) : o;
                case 'boolean':
                case 'number':
                        return o;
        }

        if (o instanceof Date) {
                return typeOrCfg === 'string' ? o.toISOString() : Math.floor(o.getTime() / 1000);
        }

        if (Array.isArray(o)) return o.map((v: any) => ToObject(v, typeOrCfg, true));

        const d: any = {};

        for (const k of Object.keys(o)) {
                const v: any = o[k];
                if (!v) continue;
                d[k] = ToObject(v, typeOrCfg[k] || {}, true);
        }

        return d;
}

// structs
// struct2ts:github.com/munsy/wrappr.ArtifactStats
export class ArtifactStats {
        downloaded: Date;
        downloaded_by: string;
        downloads: number;
        remote_downloads: number;

        constructor(data?: any) {
                const d: any = (data && typeof data === 'object') ? ToObject(data) : {};
                this.downloaded = ('downloaded' in d) ? ParseDate(d.downloaded) : new Date();
                this.downloaded_by = ('downloaded_by' in d) ? d.downloaded_by as string : '';
                this.downloads = ('downloads' in d) ? d.downloads as number : 0;
                this.remote_downloads = ('remote_downloads' in d) ? d.remote_downloads as number : 0;
        }

        toObject(): any {
                const cfg: any = {};
                cfg.downloaded = 'string';
                cfg.downloads = 'number';
                cfg.remote_downloads = 'number';
                return ToObject(this, cfg);
        }
}

// struct2ts:github.com/munsy/wrappr.ArtifactResult
export class ArtifactResult {
        repo: string;
        path: string;
        name: string;
        type: string;
        size: number;
        created: Date;
        created_by: string;
        modified: Date;
        modified_by: string;
        updated: Date;
        stats: ArtifactStats[] | null;

        constructor(data?: any) {
                const d: any = (data && typeof data === 'object') ? ToObject(data) : {};
                this.repo = ('repo' in d) ? d.repo as string : '';
                this.path = ('path' in d) ? d.path as string : '';
                this.name = ('name' in d) ? d.name as string : '';
                this.type = ('type' in d) ? d.type as string : '';
                this.size = ('size' in d) ? d.size as number : 0;
                this.created = ('created' in d) ? ParseDate(d.created) : new Date();
                this.created_by = ('created_by' in d) ? d.created_by as string : '';
                this.modified = ('modified' in d) ? ParseDate(d.modified) : new Date();
                this.modified_by = ('modified_by' in d) ? d.modified_by as string : '';
                this.updated = ('updated' in d) ? ParseDate(d.updated) : new Date();
                this.stats = Array.isArray(d.stats) ? d.stats.map((v: any) => new ArtifactStats(v)) : null;
        }

        toObject(): any {
                const cfg: any = {};
                cfg.size = 'number';
                cfg.created = 'string';
                cfg.modified = 'string';
                cfg.updated = 'string';
                return ToObject(this, cfg);
        }
}

// struct2ts:.
class gen {
        start_pos: number;
        end_pos: number;
        total: number;

        constructor(data?: any) {
                const d: any = (data && typeof data === 'object') ? ToObject(data) : {};
                this.start_pos = ('start_pos' in d) ? d.start_pos as number : 0;
                this.end_pos = ('end_pos' in d) ? d.end_pos as number : 0;
                this.total = ('total' in d) ? d.total as number : 0;
        }

        toObject(): any {
                const cfg: any = {};
                cfg.start_pos = 'number';
                cfg.end_pos = 'number';
                cfg.total = 'number';
                return ToObject(this, cfg);
        }
}

// struct2ts:github.com/munsy/wrappr.ArtifactList
export class ArtifactList {
    results: ArtifactResult[] | null;
    range: any;

    constructor(data?: any) {
        const d: any = (data && typeof data === 'object') ? ToObject(data) : {};
        this.results = Array.isArray(d.results) ? d.results.map((v: any) => new ArtifactResult(v)) : null;
        this.range = ('range' in d) ? d.range as any : {};
    }

    toObject(): any {
        const cfg: any = {};
        return ToObject(this, cfg);
    }
}

export class ArtifactoryRequest {
    public url: string;
    public repo: string;
    public username: string;
    public password: string;

    constructor() { }
}