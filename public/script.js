var app = angular.module('url-shortener', []);

app.controller('createCtrl', function($scope, $http) {

  $scope.save = function(url) {
    if (url === undefined) { return; }

    $http.post('/encode', {url: url})
        .success(function(data, status, headers, config) {
          $scope.short_url = data.url;
        })
        .error(function(data, status, headers, config) {
          alert("fail");
        });
  }
});
