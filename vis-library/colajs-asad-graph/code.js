$.getJSON("../cytostruct.json", function(json) {    

console.log(json)
   
var cy = cytoscape({
    container: document.querySelector('#cy'),
  
    boxSelectionEnabled: false,
    autounselectify: true,
  
    style: cytoscape.stylesheet()
      .selector('node')
        .css({
          'content': 'data(name)',
          'text-valign': 'center',
          'color': 'white',
          'text-outline-width': 2,
          'background-color': '#999',
          'text-outline-color': '#999'
        })
      .selector('edge')
        .css({
          'curve-style': 'bezier',
          'target-arrow-shape': 'triangle',
          'target-arrow-color': '#ccc',
          'line-color': '#ccc',
          'width': 1
        })
      .selector(':selected')
        .css({
          'background-color': 'black',
          'line-color': 'black',
          'target-arrow-color': 'black',
          'source-arrow-color': 'black'
        })
      .selector('.faded')
        .css({
          'opacity': 0.25,
          'text-opacity': 0
        })
      .selector('.service')
        .css({
          'background-color':'red',
        }),
    elements: {
        nodes: json.nodes,
        edges: json.links
    },
    layout: {
      name: 'grid',
      padding: 10
    }
  });
  
  cy.on('tap', 'node', function(e){
    console.log(e, json.servicelayer.sitea)
    if( e.name == json.servicelayer.sitea ){
      console.log("debug")
    }
  });
 
  cy.filter('node', function(e){
    console.log(e.node)
  });

});