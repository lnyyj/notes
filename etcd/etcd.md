# etcd 使用

## 包导入坑
    测试案例上导入的是 go.etcd.io/etcd/clientv3 , 但是go.etcd.io/etcd/clientv3 无法go get， 会timeout。

    这个时候会选择在github.com拉取github.com/etcd-io/etcd, 然后进行测试就会发现进入另一个坑，会报错
    错误一: 
    ```
    $GOPATH/src/go.etcd.io/etcd/clientv3/balancer/picker/err.go:25:9: cannot use &errPicker literal (type *errPicker) as type Picker in return argument:
	*errPicker does not implement Picker (wrong type for Pick method)
		have Pick(context.Context, balancer.PickInfo) (balancer.SubConn, func(balancer.DoneInfo), error)
		want Pick(balancer.PickInfo) (balancer.PickResult, error)
    $GOPATH/src/go.etcd.io/etcd/clientv3/balancer/picker/roundrobin_balanced.go:33:9: cannot use &rrBalanced literal (type *rrBalanced) as type Picker in return argument:
	*rrBalanced does not implement Picker (wrong type for Pick method)
		have Pick(context.Context, balancer.PickInfo) (balancer.SubConn, func(balancer.DoneInfo), error)
		want Pick(balancer.PickInfo) (balancer.PickResult, error)
    ```

    错误二：
    ```
    ../../../../etcd-io/etcd/clientv3/auth.go:121:72: cannot use auth.callOpts (type []"github.com/etcd-io/etcd/vendor/google.golang.org/grpc".CallOption) as type []"go.etcd.io/etcd/vendor/google.golang.org/grpc".CallOption in argument to auth.remote.AuthEnable
    ../../../../etcd-io/etcd/clientv3/auth.go:126:74: cannot use auth.callOpts (type []"github.com/etcd-io/etcd/vendor/google.golang.org/grpc".CallOption) as type []"go.etcd.io/etcd/vendor/google.golang.org/grpc".CallOption in argument to auth.remote.AuthDisable
    ../../../../etcd-io/etcd/clientv3/auth.go:131:152: cannot use auth.callOpts (type []"github.com/etcd-io/etcd/vendor/google.golang.org/grpc".CallOption) as type []"go.etcd.io/etcd/vendor/google.golang.org/grpc".CallOption in argument to auth.remote.UserAdd
    ../../../../etcd-io/etcd/clientv3/auth.go:136:144: cannot use auth.callOpts (type []"github.com/etcd-io/etcd/vendor/google.golang.org/grpc".CallOption) as type []"go.etcd.io/etcd/vendor/google.golang.org/grpc".CallOption in argument to auth.remote.UserAdd
    ```

    错误原因是直接拉取github.com/etcd-io/etcd，不是release版本，没有严格的测试或者是开发版本。 所以会报错

    这个时候去https://github.com/etcd-io/etcd/releases下面，拉取release版本，然后放在go.etcd.io/etcd/*下面就可以测试通过
