export namespace entity {
	
	export class Assignment {
	    id: number;
	    personId: number;
	    roleId: number;
	    date: string;
	
	    static createFrom(source: any = {}) {
	        return new Assignment(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.personId = source["personId"];
	        this.roleId = source["roleId"];
	        this.date = source["date"];
	    }
	}
	export class Person {
	    id: number;
	    firstName: string;
	    lastName: string;
	
	    static createFrom(source: any = {}) {
	        return new Person(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.firstName = source["firstName"];
	        this.lastName = source["lastName"];
	    }
	}
	export class PersonRole {
	    personId: number;
	    roleId: number;
	
	    static createFrom(source: any = {}) {
	        return new PersonRole(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.personId = source["personId"];
	        this.roleId = source["roleId"];
	    }
	}
	export class Role {
	    id: number;
	    name: string;
	
	    static createFrom(source: any = {}) {
	        return new Role(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	    }
	}

}

