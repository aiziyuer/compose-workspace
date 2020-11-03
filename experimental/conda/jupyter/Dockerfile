FROM conda/miniconda3-centos7

# 升级conda
RUN conda update -y -c defaults conda && conda clean -ay

RUN conda init zsh && conda init bash

# 必要工具
RUN . ~/.bashrc && conda install -y -c conda-forge python=3.6 ipython jupyterlab nodejs && conda clean -ay
RUN . ~/.bashrc && conda install -y -c conda-forge nb_conda_kernels && conda clean -ay

# 运维定位用的工具
RUN yum install -y which net-tools tcpdump && yum clean all

# 安装db连接器
RUN . ~/.bashrc && conda install -y -c anaconda mysql-connector-python && conda clean -ay

RUN . ~/.bashrc && pip install pyecharts

# 设置启动
RUN mkdir -p /data
CMD . ~/.bashrc && /usr/local/bin/jupyter lab --ip=* --port=8888 --no-browser --notebook-dir=/data --allow-root
