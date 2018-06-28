
# 给普通用户加入sudo权限
> 若想让test用户拥有sudo权限，则需要修改sudoers文件  

- 首先切换为root用户  
    su root  

- 修改sudoers文件  
    vim /etc/sudoers  

- 找到权限设置，如下  
    root    ALL=(ALL)       ALL  

- 若要给violet用户增加sudo权限，需要增加如下一行  
    root    ALL=(ALL)       ALL  
    test    ALL=(ALL)       ALL  

- 保存退出后test用户则拥有了sudo权限  
