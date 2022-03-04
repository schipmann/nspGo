/* global Promise, fetch, window, cytoscape, document, tippy, _ */

/* ::::::::::::::::::::::::::::
First Research of Document :
line 28-31 daten in den Nodes 

:::::::::::::::::::::::::::::*/
Promise.all([
  fetch('cy-style.json')
    .then(function(res) {
      return res.json();
    }),
  fetch('../cytostruct.json')
    .then(function(res) {
      return res.json();
    })
])
  .then(function(dataArray) {
    var h = function(tag, attrs, children){
      var el = document.createElement(tag);

      Object.keys(attrs).forEach(function(key){
        var val = attrs[key];
        console.log(val)

        el.setAttribute(key, val);
      });

      children.forEach(function(child){
        console.log(child)
        el.appendChild(child);
      });
      return el;
    };

    var t = function(text){
      var el = document.createTextNode(text);
      return el;
    };

    var $ = document.querySelector.bind(document);



/*     var cy = window.cy = cytoscape({
      container: document.getElementById('cy'),
      style: dataArray[0],
      elements: dataArray[1],
      layout: { name: 'random' }
    }); */

    var cy = window.cy = cytoscape({
      container: document.getElementById('cy'),
      elements: dataArray[1].nodes,
      layout: { name: 'random' }
    });
    console.log(dataArray)
    console.log(dataArray[1].nodes[0])


    var params = {
      name: 'cola',
      nodeSpacing: 5,
      nodeSpacingAsad: 10,
      edgeLengthVal: 45,
      animate: true,
      randomize: false,
      maxSimulationTime: 1500
    };
    var layout = makeLayout();

    layout.run();

    var $btnParam = h('div', {
      'class': 'param'
    }, []);

    var $config = $('#config');

    $config.appendChild( $btnParam );


    var sliders = [
      {
        label: 'Edge length',
        param: 'edgeLengthVal',
        min: 1,
        max: 200
      },

      {
        label: 'Node spacing',
        param: 'nodeSpacing',
        min: 1,
        max: 50
      },

      {
        label: 'Placeholder',
        param: 'Placeholder',
        min: 1,
        max: 500
      }
    ];

    var buttons = [
      {
        label: h('span', { 'class': 'fa fa-random' }, []),
        layoutOpts: {
          randomize: true,
          flow: null
        }
      },

      {
        label: h('span', { 'class': 'fa fa-long-arrow-down' }, []),
        layoutOpts: {
          flow: { axis: 'y', minSeparation: 30 }
        }
      }
    ];

    // button for BgpLU link
    var buttonsBgpLuLink = [
      {
        label: h('span', { 'class': 'fa fa-check-square' }, []),
        layoutOpts: {
          flow: { axis: 'y', minSeparation: 30 }
        }
      }
    ];

    sliders.forEach( makeSlider );

    buttons.forEach( makeButton );

    buttonsBgpLuLink.forEach(makeButtonBgpLuLink);


    function makeLayout( opts ){
      console.lo
      params.randomize = false;
      params.edgeLength = function(e){ return params.edgeLengthVal / e.data('weight'); };

      for( var i in opts ){
        params[i] = opts[i];
      }

      return cy.layout( params );
    }

    function makeSlider( opts ){
      var $input = h('input', {
        id: 'slider-'+opts.param,
        type: 'range',
        min: opts.min,
        max: opts.max,
        step: 1,
        value: params[ opts.param ],
        'class': 'slider'
      }, []);

      var $param = h('div', { 'class': 'param' }, []);

      var $label = h('label', { 'class': 'label label-default', for: 'slider-'+opts.param }, [ t(opts.label) ]);

      $param.appendChild( $label );
      $param.appendChild( $input );

      $config.appendChild( $param );

      var update = _.throttle(function(){
        params[ opts.param ] = $input.value;

        layout.stop();
        layout = makeLayout();
        layout.run();
      }, 1000/30);

      $input.addEventListener('input', update);
      $input.addEventListener('change', update);
    }

    function makeButton( opts ){
      var $button = h('button', { 'class': 'btn btn-default' , 'title': "Auto Layout"}, [ opts.label ]);

      $btnParam.appendChild( $button );

      $button.addEventListener('click', function(){
        layout.stop();

        if( opts.fn ){ opts.fn(); }

        layout = makeLayout( opts.layoutOpts );
        layout.run();
      });
    }


    bgpLuLink = cy.filter ('edge[group = "bgpLu"]');
    //var bgpLuLink = cy.edges();

    var bgpLuLinkVisibiliy = true

    function makeButtonBgpLuLink( opts ){
      var $buttonsBgpLuLink = h('button', { 'class': 'btn btn-default' , 'title': "Toggle show BGP-LU Link-Tunnel" }, [ opts.label ]);
      $btnParam.appendChild( $buttonsBgpLuLink );
      $buttonsBgpLuLink.addEventListener('click', function(){
        layout.stop();
        
        if (bgpLuLinkVisibiliy){
          bgpLuLink.remove()
          bgpLuLinkVisibiliy = false
          console.log('bgpLuLinkVisibiliy: ' + bgpLuLinkVisibiliy);
        }
        else {
          bgpLuLink.restore()
          bgpLuLinkVisibiliy = true
          console.log('bgpLuLinkVisibiliy: ' + bgpLuLinkVisibiliy);
        }
        
      });
    }

    console.log(cy.nodes(), "")

    var makeTippy = function(node, html){
      return tippy( node.popperRef(), {
        html: html,
        trigger: 'manual',
        arrow: true,
        placement: 'bottom',
        hideOnClick: false,
        interactive: true
      } ).tooltips[0];
    };

    var hideTippy = function(node){
      var tippy = node.data('tippy');

      if(tippy != null){
        tippy.hide();
      }
    };

    var hideAllTippies = function(){
      cy.nodes().forEach(hideTippy);
    };

    cy.on('tap', function(e){
      if(e.target === cy){
        hideAllTippies();
      }
    });

    cy.on('tap', 'edge', function(e){
      hideAllTippies();
    });

    cy.on('zoom pan', function(e){
      hideAllTippies();
    });

/*     cy.nodes().forEach(function(n){
 */      /* var g = n.data('name'); */

/*       var $links = [
        {
          name: 'System Address: ' + n.data('systemAddress'),
          url: 'http://www.genecards.org/cgi-bin/carddisp.pl?gene=' + g
        },
        {
          name:  'Router Type: ' + n.data('routerType'),
          url: 'http://www.uniprot.org/uniprot/?query='+ g +'&fil=organism%3A%22Homo+sapiens+%28Human%29+%5B9606%5D%22&sort=score'
        },
        {
          name: 'Network: ' + n.data('network'),
          url: 'https://www.linkedin.com/in/asadarafat/' + g */
      /*  }  */
/*       ].map(function( link ){
        return h('a', { target: '_blank', href: link.url, 'class': 'tip-link' }, [ t(link.name) ]);
      }); */

/*       var tippy = makeTippy(n, h('div', {}, $links));

      n.data('tippy', tippy);

      n.on('click', function(e){
        tippy.show();

        cy.nodes().not(n).forEach(hideTippy);
      });
    }); */


    $('#config-toggle').addEventListener('click', function(){
      $('body').classList.toggle('config-closed');

      cy.resize();
    });

  });
