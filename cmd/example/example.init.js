window.addEventListener("load", function load(event){

    var select_el = document.getElementById("placetypes");
    var ancestors_el = document.getElementById("ancestors");
    var descendants_el = document.getElementById("descendants");    
    var parents_el = document.getElementById("parents");
    var children_el = document.getElementById("children");    

    var set_parents = function(pt){
	
	whosonfirst_placetypes_parents(pt)
	    .then((data) => {
		
		var placetypes = JSON.parse(data);
		var count = placetypes.length;
		
		for (var i=0; i < count; i++) {
		    
		    var pt = placetypes[i];
		    var name = pt["name"];
		    
		    var item = document.createElement("li");
		    item.appendChild(document.createTextNode(name));
		    parents_el.appendChild(item);
		}
		
	    }).catch((err) => {
		console.log("SAD", err);
	    });
	
    };

    var set_children = function(pt){
	
	whosonfirst_placetypes_children(pt)
	    .then((data) => {
		
		var placetypes = JSON.parse(data);
		var count = placetypes.length;
		
		for (var i=0; i < count; i++) {
		    
		    var pt = placetypes[i];
		    var name = pt["name"];
		    
		    var item = document.createElement("li");
		    item.appendChild(document.createTextNode(name));
		    children_el.appendChild(item);
		}
		
	    }).catch((err) => {
		console.log("SAD", err);
	    });
	
    };
    
    var set_descendants = function(pt){
	
	whosonfirst_placetypes_descendants(pt, "common,optional,common_optional")
	    .then((data) => {
		
		var placetypes = JSON.parse(data);
		var count = placetypes.length;
		
		for (var i=0; i < count; i++) {
		    
		    var pt = placetypes[i];
		    var name = pt["name"];
		    
		    var item = document.createElement("li");
		    item.appendChild(document.createTextNode(name));
		    descendants_el.appendChild(item);
		}
		
	    }).catch((err) => {
		console.log("SAD", err);
	    });
	
    };

    var set_ancestors = function(pt){
	
	whosonfirst_placetypes_ancestors(pt, "common,optional,common_optional")
	    .then((data) => {
		
		var placetypes = JSON.parse(data);
		var count = placetypes.length;
		
		for (var i=0; i < count; i++) {
		    
		    var pt = placetypes[i];
		    var name = pt["name"];
		    
		    var item = document.createElement("li");
		    item.appendChild(document.createTextNode(name));
		    ancestors_el.appendChild(item);
		}
		
	    }).catch((err) => {
		console.log("SAD", err);
	    });
    };

    // https://github.com/sfomuseum/go-http-wasm
    // https://github.com/sfomuseum/go-http-wasm/blob/main/static/javascript/sfomuseum.wasm.js
    
    sfomuseum.wasm.fetch("wasm/whosonfirst_placetypes.wasm").then(rsp => {
	
	whosonfirst_placetypes()
	    .then((data) => {

		var names = [];
		
		var placetypes = JSON.parse(data);
		var count = placetypes.length;
		
		for (var i=0; i < count; i++){
		    var pt = placetypes[i];
		    var name = pt["name"];
		    names.push(name);
		}

		names.sort()

		for (var i=0; i < count; i++){

		    var name = names[i];
		    
		    var opt = document.createElement("option");
		    opt.setAttribute("value", name);
		    opt.appendChild(document.createTextNode(name));

		    select_el.appendChild(opt);
		}
		
		select_el.onchange = function(){

		    ancestors_el.innerHTML = "";
		    descendants_el.innerHTML = "";
		    parents_el.innerHTML = "";
		    children_el.innerHTML = "";
		    
		    var pt = select_el.value;

		    if (pt == ""){
			return;
		    }

		    set_descendants(pt);
		    set_ancestors(pt);
		    set_children(pt);
		    set_parents(pt);
		};
		
	    })
	    .catch ((err) => {
		console.log("SAD", err);
	    });
	
    });
    
});
